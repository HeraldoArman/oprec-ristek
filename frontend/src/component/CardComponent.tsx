import { Card, Dropdown } from "flowbite-react";

interface CardComponentProps {
  title: string;
  description: string;
  image: string;
  buttonLabel: string;
  buttonStyle: string;
  addTryout?: boolean;
}

const CardComponent = ({
  title,
  description,
  image,
  buttonLabel,
  buttonStyle,
  addTryout = false,
}: CardComponentProps) => {
  return (
    <Card className="relative w-full p-4 shadow-md">
      {addTryout && (
        <div className="absolute top-2 right-2">
          <Dropdown inline label="">
            <Dropdown.Item>
              <a
                href="#"
                className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-600 dark:hover:text-white"
              >
                Edit
              </a>
            </Dropdown.Item>
            <Dropdown.Item></Dropdown.Item>
            <Dropdown.Item>
              <a
                href="#"
                className="block px-4 py-2 text-sm text-red-600 hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-600 dark:hover:text-white"
              >
                Delete
              </a>
            </Dropdown.Item>
          </Dropdown>
        </div>
      )}
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
