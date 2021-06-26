import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";

const IndexPage = () => {
  return (
    <FullPage>
      <Header title="Register" />
      <Responsive align="center" justify="center">
        <div>Hello World</div>
      </Responsive>
    </FullPage>
  );
};

export default IndexPage;
