import React, { useEffect, useState } from "react";
import { NextPage } from "next";
import RegisterForm from "../src/components/layout/RegisterForm";
import { UserProfile } from "../src/user";
import { getProfile, updateProfile } from "../src/network";

const Page: NextPage = () => {
  const [user, setUser] = useState<UserProfile>();
  const [settingComplete, setSettingComplete] = useState(false);
  useEffect(() => {
    getProfile().then((u) => setUser(u));
  }, []);
  const onSubmit = (user: UserProfile) => {
    updateProfile(user)
      .then((u) => {
        console.log(u);
        setSettingComplete(true);
      })
      .catch((err) => console.error(err));
  };
  return <>{user ? <RegisterForm formName="会員情報設定" user={user} onSubmit={onSubmit} /> : "Loading..."}</>;
};

export default Page;
