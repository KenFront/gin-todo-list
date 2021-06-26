import { useRouter } from "next/router";

export const useChangePage = () => {
  const router = useRouter();

  return {
    changePath: (path: string) => {
      if (router.pathname != path) {
        router.push(path);
      }
    },
  };
};
