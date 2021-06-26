import { Heading } from "@chakra-ui/react";

import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";

const ErrorPage = () => {
  return (
    <FullPage>
      <Header title="Error" />
      <Responsive align="center" justify="center">
        <Heading size="lg">Page not found</Heading>
      </Responsive>
    </FullPage>
  );
};

export default ErrorPage;
