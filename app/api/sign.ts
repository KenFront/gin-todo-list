import { ClientRequest, PostOptions } from "@/lib/request";
import { ApiOnSuccess } from "@/lib/request/type";

export const signIn = async (val: { account: string; password: string }) => {
  const res = await ClientRequest({
    path: `/signin`,
    options: PostOptions(val),
  });

  return res as ApiOnSuccess<string>;
};

export const signOut = async () => {
  const res = await ClientRequest({
    path: "/signout",
    options: PostOptions()
  });

  return res as ApiOnSuccess<string>;
};
