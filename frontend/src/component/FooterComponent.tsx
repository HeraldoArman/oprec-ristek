import { Footer } from "flowbite-react";
import { BsGithub, BsLinkedin } from "react-icons/bs";

export function FooterComponent() {
  return (
    <Footer container>
      <div className="w-full">
        <div className="grid w-full justify-between sm:flex sm:justify-between md:flex md:grid-cols-1"></div>
        <Footer.Divider />
        <div className="w-full sm:flex sm:items-center sm:justify-between">
          <Footer.Copyright href="#" by="Aldo" year={2025} />
          <div className="mt-4 flex space-x-6 sm:mt-0 sm:justify-center">
            <Footer.Icon
              href="https://www.linkedin.com/in/heraldo-arman/"
              icon={BsLinkedin}
            />
            <Footer.Icon
              href="https://github.com/HeraldoArman"
              icon={BsGithub}
            />
          </div>
        </div>
      </div>
    </Footer>
  );
}
export default FooterComponent;
