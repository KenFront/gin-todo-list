import { useRouter } from "next/router";
import { compile } from "path-to-regexp";

export const useChangePage = () => {
  const router = useRouter();

  return {
    changePath: ({
      path,
      param,
    }: {
      path: string;
      param?: { [key: string]: string };
    }) => {
      let realPath: string;
      if (param) {
        const toPath = compile(path);
        realPath = toPath(param);
      } else {
        realPath = path;
      }
      if (router.pathname != realPath) {
        router.push(realPath);
      }
    },
  };
};
