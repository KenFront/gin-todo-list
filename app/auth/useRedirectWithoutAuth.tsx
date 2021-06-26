import { ROUTE } from "@/route";
import { useChangePage } from "@/lib/hook/useChangePage";

export const useRedirectWithoutAuth = () => {
  const { changePath } = useChangePage();
  return {
      redirect: () => {
        changePath(ROUTE.SIGN_IN)
      }
  }
};
