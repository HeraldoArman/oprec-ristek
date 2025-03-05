import { Flowbite } from "flowbite-react";
import Navbar from "./component/NavbarComponent";
// import CardComponent from "./component/CardComponent";
import BoxComponent from "./component/BoxComponent";
import FooterComponent from "./component/FooterComponent";
// import FloatingButton from "./component/FloatingButton";
const tryoutList = [
  {
    title: "kuis DDP",
    description : "quiz nya ez kok cuma hello world",
    image: "https://picsum.photos/1000/600",
    buttonLabel: "Start Tryout",
    buttonStyle: "bg-green-500 hover:bg-green-600",
  },
  {
    title: "kuis sesuatu",
    description : "quiz apaan yak",

    image: "https://picsum.photos/1000/600",

    buttonLabel: "Tryout closed",
    buttonStyle: "bg-gray-300 cursor-not-allowed",
  },
  {
    title: "kuis kalkulus",
    description : "muehehehehehehe kalkulus susah banget loh",

    image: "https://picsum.photos/1000/600",

    buttonLabel: "Tryout closed",
    buttonStyle: "bg-gray-300 cursor-not-allowed",
  },
];


function App() {
  return (
    <Flowbite>
      <div className="min-h-screen bg-gray-100">

        <Navbar />

        <div className="container mx-auto px-5 py-5">
          <BoxComponent name={"My Tryout"} tryout={tryoutList} addTryout={true}/>
        </div>
        <div className="container mx-auto px-5 py-5">
          <BoxComponent name={"All Tryout"} tryout={[...tryoutList, ...tryoutList]} addTryout={false}/>
        </div>
        <FooterComponent/>
      </div>
      
    </Flowbite>
  );
}

export default App;
