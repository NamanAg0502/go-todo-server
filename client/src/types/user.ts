type User = {
  id: string;
  name: string;
  email: string;
  password: string;
  createdAt: Date;
  updatedAt: Date;
};

type ApiResponse = {
  message: string;
  data: any;
  success: boolean;
};
