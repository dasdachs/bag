export interface Item {
  id: number;
  itemName: string;
  quantity: number;
  categoryId: number;
  category: {
    id: number;
    categoryName: string;
    items: unknown[];
  };
}
