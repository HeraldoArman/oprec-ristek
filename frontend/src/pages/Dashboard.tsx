import { useEffect, useState } from "react";
import { Flowbite, Toast } from "flowbite-react";
import { motion, AnimatePresence } from "framer-motion"; // Import Framer Motion
import { HiX } from "react-icons/hi";
import NavbarComponent from "../component/NavbarComponent";
import BoxComponent from "../component/BoxComponent";
import FooterComponent from "../component/FooterComponent";
// import { useNavigate } from "react-router-dom";


interface Tryout {
  id: number;
  title: string;
  detail: string;
  image?: string;
}

function Dashboard() {
  // const navigate = useNavigate();

  const [userTryouts, setUserTryouts] = useState<Tryout[]>([]);
  const [allTryouts, setAllTryouts] = useState<Tryout[]>([]);
  const [loadingUser, setLoadingUser] = useState(true);
  const [loadingAll, setLoadingAll] = useState(true);
  const [errorUser, setErrorUser] = useState("");
  const [errorAll, setErrorAll] = useState("");
  const [showToast, setShowToast] = useState(false);

  useEffect(() => {
    const fetchUserTryout = async () => {
      try {
        const response = await fetch("http://127.0.0.1:3000/tryout/user/john_doe3");
        if (!response.ok) throw new Error("Gagal mengambil data My Tryout");
        const data = await response.json();
        setUserTryouts(data.map((item: any) => ({
          id: item.ID,
          title: item.title,
          detail: item.detail,
          image: "https://picsum.photos/1000/600",
        })));
      } catch (err: any) {
        setErrorUser(err.message);
      } finally {
        setLoadingUser(false);
      }
    };

    const fetchAllTryout = async () => {
      try {
        const response = await fetch("http://127.0.0.1:3000/tryout/");
        if (!response.ok) throw new Error("Gagal mengambil data All Tryout");
        const data = await response.json();
        setAllTryouts(data.map((item: any) => ({
          id: item.ID,
          title: item.title,
          detail: item.detail,
          image: "https://picsum.photos/1000/600",
        })));
      } catch (err: any) {
        setErrorAll(err.message);
      } finally {
        setLoadingAll(false);
      }
    };

    fetchUserTryout();
    fetchAllTryout();
  }, []);

  // Fungsi untuk menghapus tryout & menampilkan toast
  const handleDeleteTryout = (id: number) => {
    setUserTryouts((prev) => prev.filter((tryout) => tryout.id !== id));
    setShowToast(true); // Tampilkan toast
    setTimeout(() => setShowToast(false), 3000); // Sembunyikan setelah 3 detik
  };

  return (
    <Flowbite>
      <div className="flex min-h-screen flex-col bg-gray-100 relative">
        <NavbarComponent />

        <div className="container mx-auto flex-grow px-5 py-5">
          <BoxComponent
            name="My Tryout"
            addTryout={true}
            tryouts={userTryouts}
            loading={loadingUser}
            error={errorUser}
            setTryouts={handleDeleteTryout}
          />
        </div>

        <div className="container mx-auto px-5 py-5">
          <BoxComponent
            name="All Tryout"
            addTryout={false}
            tryouts={allTryouts}
            loading={loadingAll}
            error={errorAll}
            setTryouts={handleDeleteTryout}
          />
        </div>

        <FooterComponent />

        {/* Toast dengan animasi */}
        <div className="fixed top-5 left-5 z-50">
          <AnimatePresence>
            {showToast && (
              <motion.div
                initial={{ opacity: 0, y: -20 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: -20 }}
                transition={{ duration: 0.3 }}
              >
                <Toast>
                  <div className="inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-red-100 text-red-500 dark:bg-red-800 dark:text-red-200">
                    <HiX className="h-5 w-5" />
                  </div>
                  <div className="ml-3 text-sm font-normal">Item has been deleted.</div>
                  <Toast.Toggle onClick={() => setShowToast(false)} />
                </Toast>
              </motion.div>
            )}
          </AnimatePresence>
        </div>
      </div>
    </Flowbite>
  );
}

export default Dashboard;
