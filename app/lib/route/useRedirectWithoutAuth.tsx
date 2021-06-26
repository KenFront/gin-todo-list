import { ROUTE } from "./index";
import { useChangePage } from "../hook/useChangePage";

export const useRedirectWithoutAuth = () => {
  const { changePath } = useChangePage();
  return {
      redirect: () => {
        changePath(ROUTE.SIGN_IN)
      }
  }
};
