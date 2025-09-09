export type ConnectionState = "connected" | "connecting" | "disconnected";

export type LogEntry = {
  source: string;
  data: string;
};

export type ResourceMetrics = {
  cpu: string;
  memory: string;
  gpu: string;
  gpuMemory: string;
};
