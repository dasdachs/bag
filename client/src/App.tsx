import React, { useState, useEffect } from "react";

import { Client as Styletron } from "styletron-engine-atomic";
import { Provider as StyletronProvider } from "styletron-react";
import { BaseProvider, LightTheme } from "baseui";
import { Table } from "baseui/table-semantic";
import { Option } from "baseui/select";

import { AddItem } from "./containers/add-item";

export interface Item {
  id: number;
  name: string;
  quantity: number;
  categoryId: number;
  category: {
    id: number;
    name: string;
    items: unknown[];
  };
}

export interface Category {
  readonly id: number;
  readonly name: string;
}

const engine = new Styletron();

export interface CategoryOption extends Option {
  readonly category: Category;
}

function App(): JSX.Element {
  const columns = ["Id", "Name", "Quantity", "Category"];
  const [items, setItems] = useState<(string | number)[][]>([]);
  const [categories, setCategories] = useState<CategoryOption[]>([]);

  useEffect(() => {
    (async () => {
      const request = await fetch("/api/v1/categories");

      if (request.ok) {
        const categories: Category[] = await request.json();
        const categoryOptions = categories.map((category) => ({
          id: category.id,
          label: category.name,
          category,
        }));
        setCategories(categoryOptions);
      }
    })();
  }, []);

  useEffect(() => {
    (async () => {
      const request = await fetch("/api/v1/items");

      if (request.ok) {
        const items: Item[] = await request.json();
        const columns = items.map((item) => [
          item.id,
          item.name,
          item.quantity,
          item.category.name,
        ]);
        setItems(columns);
      }
    })();
  }, []);

  return (
    <StyletronProvider value={engine}>
      <BaseProvider theme={LightTheme}>
        <header>
          <h1>Inventory</h1>
        </header>
        <AddItem categories={categories} items={items} />
        <Table columns={columns} data={items} />
      </BaseProvider>
    </StyletronProvider>
  );
}

export default App;
