import { useState, useRef, useCallback, DependencyList } from 'react'

export type AsyncStatus = 'idle' | 'loading' | 'success' | 'error'

export const useAsync = <T, K extends Array<unknown>>(
  fn: (...arg: K) => Promise<T>,
  deps: DependencyList = []
) => {
  const [result, setResult] = useState<{
    status: AsyncStatus
    success: T | undefined
    error: Error | undefined
  }>({
    status: 'idle',
    success: undefined,
    error: undefined
  })
  const lastRunningId = useRef(0)

  const reset = useCallback(() => {
    if (result.status !== 'loading') {
      setResult({
        status: 'idle',
        success: undefined,
        error: undefined
      })
    }
  }, [result])

  const execute = useCallback(async (...arg: K) => {
    const runningId = ++lastRunningId.current
    try {
      setResult({
        status: 'loading',
        success: undefined,
        error: undefined
      })
      const result = await fn(...arg)
      if (runningId === lastRunningId.current) {
        setResult({
          status: 'success',
          success: result,
          error: undefined
        })
      }
    } catch (error) {
      if (runningId === lastRunningId.current) {
        setResult({
          status: 'error',
          success: undefined,
          error: error
        })
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, deps)

  return {
    status: result.status,
    result: result.success,
    error: result.error,
    execute,
    reset
  }
}
