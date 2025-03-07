import { Avatar, Dropdown, Navbar } from "flowbite-react";
// import { useNavigate } from "react-router-dom";

const NavbarComponent = () => {
  // const navigate = useNavigate();
  return (
    <Navbar fluid rounded>
      <Navbar.Brand href={"/"}>
        <span className="self-center whitespace-nowrap text-xl font-bold text-green-500">
          RIZZ QUIZ
        </span>
      </Navbar.Brand>
      <div className="flex md:order-2">
        <Dropdown arrowIcon={false} inline label={<Avatar rounded />}>
          <Dropdown.Header>
            <span className="block text-sm">John Doe</span>
            <span className="block truncate text-sm font-medium">
              johndoe@ristek.ui.ac.id
            </span>
          </Dropdown.Header>
          <Dropdown.Item>Dashboard</Dropdown.Item>
          <Dropdown.Divider />
          <Dropdown.Item>Sign out</Dropdown.Item>
        </Dropdown>
      </div>
    </Navbar>
  );
};
export default NavbarComponent;
