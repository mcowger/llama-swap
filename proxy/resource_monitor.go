package proxy

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type ResourceMetrics struct {
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
	GPU       string `json:"gpu"`
	GPUMemory string `json:"gpuMemory"`
}

type ResourceMonitor struct {
	config  ResourceMonitorConfig
	metrics ResourceMetrics
	mutex   sync.RWMutex
	stopCh  chan struct{}
	logger  *slog.Logger
}

func NewResourceMonitor(config ResourceMonitorConfig, logger *slog.Logger) *ResourceMonitor {
	return &ResourceMonitor{
		config: config,
		logger: logger.With("component", "ResourceMonitor"),
		stopCh: make(chan struct{}),
	}
}

func (rm *ResourceMonitor) hasCommands() bool {
	return rm.config.CPU.Command != "" ||
		rm.config.Memory.Command != "" ||
		rm.config.GPU.Command != "" ||
		rm.config.GPUMemory.Command != ""
}

func (rm *ResourceMonitor) Start() {
	if rm.config.Interval <= 0 || !rm.hasCommands() {
		return
	}

	rm.logger.Info("Starting resource monitor")
	ticker := time.NewTicker(time.Duration(rm.config.Interval) * time.Second)
	go func() {
		rm.updateMetrics()
		for {
			select {
			case <-ticker.C:
				rm.updateMetrics()
			case <-rm.stopCh:
				ticker.Stop()
				return
			}
		}
	}()
}

func (rm *ResourceMonitor) Stop() {
	rm.logger.Info("Stopping resource monitor")
	close(rm.stopCh)
}

func (rm *ResourceMonitor) GetMetrics() ResourceMetrics {
	rm.mutex.RLock()
	defer rm.mutex.RUnlock()
	return rm.metrics
}

func (rm *ResourceMonitor) updateMetrics() {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	if rm.config.CPU.Command != "" {
		rm.metrics.CPU = rm.executeCommand(rm.config.CPU.Command) + rm.config.CPU.Units
	}
	if rm.config.Memory.Command != "" {
		rm.metrics.Memory = rm.executeCommand(rm.config.Memory.Command) + rm.config.Memory.Units
	}
	if rm.config.GPU.Command != "" {
		rm.metrics.GPU = rm.executeCommand(rm.config.GPU.Command) + rm.config.GPU.Units
	}
	if rm.config.GPUMemory.Command != "" {
		rm.metrics.GPUMemory = rm.executeCommand(rm.config.GPUMemory.Command) + rm.config.GPUMemory.Units
	}
}

func (rm *ResourceMonitor) executeCommand(command string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	out, err := exec.CommandContext(ctx, "bash", "-c", command).Output()
	if err != nil {
		rm.logger.Error("Error executing command", "command", command, "error", err)
		return "error"
	}
	return strings.TrimSpace(string(out))
}
