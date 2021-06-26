import { ROUTE } from "./index";
import { useChangePage } from "../hook/useChangePage";

export const useRedirectWithAuth = () => {
  const { changePath } = useChangePage();
  return {
      redirect: () => {
        changePath(ROUTE.INDEX)
      }
  }
};
