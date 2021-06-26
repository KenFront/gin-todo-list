import { Button } from "@chakra-ui/react";

import { signOut } from "@/api/sign";
import { useRedirectWithoutAuth } from "./useRedirectWithoutAuth";

export const SignOutButton = () => {
  const { redirect } = useRedirectWithoutAuth();
  return (
    <Button
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
  );
};
