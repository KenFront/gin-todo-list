import { ROUTE } from "@/route";
import { useChangePage } from "@/lib/hook/useChangePage";

export const useRedirectWithAuth = () => {
  const { changePath } = useChangePage();
  return {
    redirect: () => {
      changePath({ path: ROUTE.INDEX });
    },
  };
};
