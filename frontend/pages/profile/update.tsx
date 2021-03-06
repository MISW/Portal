import React from "react";
import { NextPage } from "next";
import { Alert } from "@material-ui/lab";
import RegisterForm, {
  SubmitResult,
} from "../../src/components/layout/RegisterForm";
import { ConfigurableProfile } from "../../src/user";
import { withLogin } from "../../src/middlewares/withLogin";
import { selectCurrentUser, updateCurrentUser } from "features/currentUser";
import { nonNullOrThrow } from "utils";
import { unwrapResult } from "@reduxjs/toolkit";
import { useSelector, useDispatch } from "react-redux";

const Page: NextPage = () => {
  const dispatch = useDispatch();
  const currentUser = nonNullOrThrow(useSelector(selectCurrentUser));
  const onSubmit = async (user: ConfigurableProfile): Promise<SubmitResult> => {
    try {
      await dispatch(updateCurrentUser(user)).then(unwrapResult);
      return { status: "success" as const };
    } catch (e) {
      console.error(e);
      return { status: "error", message: "エラーが発生しました" };
    }
  };
  return (
    <RegisterForm
      formName="会員情報設定"
      formType="setting"
      user={currentUser}
      onSubmit={onSubmit}
      successMessage={
        <Alert severity="success">プロフィールが変更されました!</Alert>
      }
    />
  );
};

export default withLogin(Page);
