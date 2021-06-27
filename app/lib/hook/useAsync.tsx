import { useState, useRef, useCallback, DependencyList } from "react";

export type AsyncStatus = "idle" | "loading" | "success" | "error";

export const useAsync = <T, K extends Array<unknown>>(
  fn: (...arg: K) => Promise<T>,
  deps: DependencyList = []
) => {
  const [status, setStatus] = useState<AsyncStatus>("idle");
  const [result, setResult] = useState<T>();
  const [error, setError] = useState<Error>();
  const lastRunningId = useRef(0);

  const reset = useCallback(() => {
    if (status !== "idle") {
      setStatus("idle");
      setResult(undefined);
      setError(undefined);
    }
  }, [status]);

  const execute = useCallback(async (...arg: K) => {
    const runningId = ++lastRunningId.current;
    try {
      reset();
      setStatus("loading");
      const res = await fn(...arg);
      if (runningId === lastRunningId.current) {
        setResult(res);
        setStatus("success");
      }
    } catch (error) {
      if (runningId === lastRunningId.current) {
        setError(error);
        setStatus("error");
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, deps);

  return {
    status,
    result,
    error,
    execute,
    reset,
  };
};
