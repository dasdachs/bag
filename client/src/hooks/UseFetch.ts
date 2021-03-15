import { useEffect, useState } from "react";

interface UseFetchReturn<T> {
  abortController: AbortController;
  error: Error | null;
  isLoading: boolean;
  data: T[];
}

export function UseFetch<T>(
  url: string,
  init?: RequestInit
): UseFetchReturn<T> {
  const abortController = new AbortController();

  const [data, setData] = useState<T[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const fetchItems = async () => {
      setIsLoading(true);
      try {
        const resp = await fetch(url, init);

        if (resp.ok) {
          try {
            setData(await resp.json());
          } catch (e) {
            setError(e as Error);
          }
        } else {
          const error = new Error("Could not fetch data");
          error.name = "Response Error";
          error.stack = `${resp.url} ${resp.statusText} \n${resp.body}`;
          setError(error);
        }
      } catch (e) {
        setError(e as Error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchItems();
  }, [url, init]);

  return {
    abortController,
    data,
    error,
    isLoading,
  };
}
