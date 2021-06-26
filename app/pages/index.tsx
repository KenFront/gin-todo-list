import { Button } from "@chakra-ui/react";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";
import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";
import { signOut } from "@/api/sign";
import { useRedirectWithoutAuth } from "@/auth/useRedirectWithoutAuth";

export const getServerSideProps = CheckPageWithAuth;

const IndexPage = () => {
  const { redirect } = useRedirectWithoutAuth();
  return (
    <FullPage>
      <Header title="Index" />
      <Responsive>
        <Button
          mt={4}
          colorScheme="teal"
          onClick={async () => {
            try {
              await signOut();
              redirect();
            } catch (e) {
              console.error(e);
            }
          }}
        >
          Sign out
        </Button>
      </Responsive>
    </FullPage>
  );
};

export default IndexPage;
