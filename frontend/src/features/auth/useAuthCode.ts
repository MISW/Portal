import { useCallback, useState } from "react";
import { useHTTPClient } from "infra/api";
import { loadState } from "./restoreState";

export const useAuthCode = () => {
  const [error, setError] = useState<unknown>();
  const http = useHTTPClient();
  const processAuthCode = useCallback(
    async (code: string, state: string) => {
      try {
        await http.post("public/callback", { json: { code, state } }).json();
        return loadState();
      } catch (e) {
        setError(e);
      }
    },
    [http]
  );

  return {
    processAuthCode,
    error,
  } as const;
};