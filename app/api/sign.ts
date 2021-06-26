import { Request, PostOptions } from "@/lib/request";
import { ApiOnSuccess } from "@/lib/request/type";

export const signIn = async (val: { account: string; password: string }) => {
  const res = await Request({
    path: "/api/signin",
    options: PostOptions(val),
  });

  return res as ApiOnSuccess<string>;
};

export const signOut = async () => {
  const res = await Request({
    path: "/api/signout",
    options: PostOptions()
  });

  return res as ApiOnSuccess<string>;
};