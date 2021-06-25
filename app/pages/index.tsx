import { Button } from "@chakra-ui/react";

import { CheckPageWithAuth } from "@/lib/auth/CheckPageWithAuth";
import { FullPage } from "@/lib/component/FullPage";
import { useAppToast } from "@/lib/hook/useAppToast";

export const getServerSideProps = CheckPageWithAuth;

const signOut = async () => {
  const res = await fetch("/api/signout", {
    method: "POST",
  }).then((res) => res.json());
  return res;
};

const IndexPage = () => {
  const { toastSuccess } = useAppToast();
  return (
    <FullPage>
      <Button
        mt={4}
        colorScheme="teal"
        onClick={async () => {
          const res = await signOut();
          toastSuccess({
            title: "Success",
            description: res.data,
            onCloseComplete: () => {
              window.location.reload();
            },
          });
        }}
      >
        Sign out
      </Button>
    </FullPage>
  );
};

export default IndexPage;
