import { Button } from "@chakra-ui/react";

import { CheckPageWithAuth } from "@/lib/auth/CheckPageWithAuth";
import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";
import { signOut } from "@/lib/API/sign";

export const getServerSideProps = CheckPageWithAuth;

const IndexPage = () => {
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
              window.location.reload();
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
