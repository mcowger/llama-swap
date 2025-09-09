import { useState, useEffect } from "react";
import type { ResourceMetrics } from "../lib/types";

const ResourceMonitor = () => {
  const [metrics, setMetrics] = useState<ResourceMetrics | null>(null);

  useEffect(() => {
    const fetchMetrics = () => {
      fetch("/api/resources")
        .then((res) => res.json())
        .then((data) => setMetrics(data))
        .catch((err) => console.error("Error fetching resource metrics:", err));
    };

    fetchMetrics();
    const interval = setInterval(fetchMetrics, 5000);

    return () => clearInterval(interval);
  }, []);

  if (!metrics) {
    return null;
  }

  return (
    <div className="flex items-center space-x-4 text-xs text-gray-500">
      {metrics.cpu && <span>CPU: {metrics.cpu}</span>}
      {metrics.memory && <span>Memory: {metrics.memory}</span>}
      {metrics.gpu && <span>GPU: {metrics.gpu}</span>}
      {metrics.gpuMemory && <span>GPU Memory: {metrics.gpuMemory}</span>}
    </div>
  );
};

export default ResourceMonitor;
