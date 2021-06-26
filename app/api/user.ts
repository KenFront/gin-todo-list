import { Request, PostOptions } from "@/lib/request";
import { ApiOnSuccess } from "@/lib/request/type";

export const register = async (val: {
  name: string;
  account: string;
  password: string;
  email: string;
}) => {
  const res = await Request({
    path: "/api/users",
    options: PostOptions(val),
  });

  return res as ApiOnSuccess<unknown>;
};
