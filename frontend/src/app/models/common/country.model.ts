export interface Country {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
  name: string;
  niceName: string;
  iso: string;
  iso3: string;
  numCode: number;
  phoneCode: number;
}
