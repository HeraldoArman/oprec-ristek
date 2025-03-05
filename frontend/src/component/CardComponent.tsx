import { Card, Progress } from "flowbite-react";

interface CardComponentProps {
  title: string; description : string; image: string; buttonLabel: string; buttonStyle: string;
}

const CardComponent = ({
  title,
  description,
  image,
  buttonLabel,
  buttonStyle,
}: CardComponentProps) => {
  return (
    <Card className="w-full p-4 shadow-md">
      <img
        src={image}
        alt={title}
        className="h-40 w-full rounded-t-lg object-cover"
      />
      <div className="p-4">
        <h3 className="text-center text-lg font-semibold">{title}</h3>
        <p className="text-center text-sm text-gray-600 pb-2">{description}</p>
        <button className={`w-full rounded-lg py-2 text-white ${buttonStyle}`}>
          {buttonLabel}
        </button>
      </div>
    </Card>
  );
};
export default CardComponent;
