import React, { useEffect } from "../src_furui/react";
import { NextPage } from "../src_furui/next";
import { Typography } from "../src_furui/@material-ui/core";
import { useAuthCallback } from "../src_furui/features/auth";

const Page: NextPage = () => {
  const { handleCallback } = useAuthCallback();
  useEffect(() => {
    const params = new URLSearchParams(location.search);
    const code = params.get("code");
    const state = params.get("state");
    if (code == null || state == null) {
      throw new Error("There is no status and code in query parameter");
    }
    handleCallback(code, state);
  }, [handleCallback]);
  return (
    <>
      <Typography>Loading...</Typography>
    </>
  );
};

export default Page;