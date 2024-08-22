import { LoaderPinwheel } from 'lucide-react';

const Loading = () => {
  return (
    <div className="min-h-screen w-full bg-background flex items-center justify-center">
      <LoaderPinwheel className="text-foreground" />
    </div>
  );
};

export default Loading;
