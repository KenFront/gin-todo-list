import { useState, useRef, useCallback, DependencyList } from "react";

export type AsyncStatus = "idle" | "loading" | "success" | "error";

type State<T, K extends Array<unknown>> =
  | {
      status: "idle";
      result: null;
      error: null;
      execute: (...arg: K) => Promise<void>;
      reset: () => void;
    }
  | {
      status: "loading";
      result: null;
      error: null;
      execute: (...arg: K) => Promise<void>;
      reset: () => void;
    }
  | {
      status: "success";
      result: T;
      error: null;
      execute: (...arg: K) => Promise<void>;
      reset: () => void;
    }
  | {
      status: "error";
      result: null;
      error: Error;
      execute: (...arg: K) => Promise<void>;
      reset: () => void;
    };

export const useAsync = <T, K extends Array<unknown>>(
  fn: (...arg: K) => Promise<T>,
  deps: DependencyList = []
): State<T, K> => {
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
        setStatus("error");
        setError(error);
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, deps);

  if (status === "idle") {
    return {
      status,
      result: null,
      error: null,
      execute,
      reset,
    };
  }

  if (typeof result != "undefined") {
    return {
      status: "success",
      result,
      error: null,
      execute,
      reset,
    };
  }

  if (typeof error != "undefined") {
    return {
      status: "error",
      result: null,
      error,
      execute,
      reset,
    };
  }

  return {
    status: "loading",
    result: null,
    error: null,
    execute,
    reset,
  };
};
