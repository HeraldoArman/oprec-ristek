import { Routes, Route } from "react-router-dom";
import Dashboard from "../pages/Dashboard";
import CreateTryout from "../pages/CreateTryout";
import EditTryout from "../pages/EditTryout";
import TryoutDetail from "../pages/TryoutDetail";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      {/* <Route path="/tryout/:id/edit" element={<EditTryout />} /> */}
      <Route path="/create" element={<CreateTryout />} />
      <Route path="/edit/:id" element={<EditTryout />} />
      <Route path="/detail/:id" element={<TryoutDetail />} />
    </Routes>
  );
};

export default AppRoutes;
