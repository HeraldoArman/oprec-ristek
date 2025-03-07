import { Routes, Route } from "react-router-dom";
import Dashboard from "../pages/Dashboard";
import CreateTryout from "../pages/CreateTryout";
// import EditTryout from "../pages/tryout/EditTryout";
import EditTryout from "../pages/EditTryout";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      {/* <Route path="/tryout/:id/edit" element={<EditTryout />} /> */}
      <Route path="/create" element={<CreateTryout/>} />
      <Route path="/edit/:id" element={<EditTryout />} />
    </Routes>
  );
};

export default AppRoutes;
