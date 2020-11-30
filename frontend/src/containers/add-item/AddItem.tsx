import React, { useState } from "react";

import { FormControl } from "baseui/form-control";
import { Button, SIZE as ButtonSize, SHAPE as ButtonType } from "baseui/button";
import { Input } from "baseui/input";
import { Slider } from "baseui/slider";
import {
  Select,
  SIZE as SelectSize,
  TYPE as SelectType,
  OnChangeParams,
} from "baseui/select";

import { Category, CategoryOption, Item } from "../../App";

interface AddItemProps {
  categories: CategoryOption[];
  items: (string | number)[][];
}

export const AddItem = ({ categories, items }: AddItemProps): JSX.Element => {
  const [name, setName] = useState("");
  const [quantity, setQuantity] = useState([0]);
  const [category, setCategory] = useState<CategoryOption[]>([]);

  const onSelect = async (param: OnChangeParams): Promise<void> => {
    let categoryOption = categories.find(
      (category) => category.id === param?.option?.id ?? -1
    );
    if (!categoryOption) {
      const response = await fetch("/api/v1/categories", {
        method: "POST",
        body: JSON.stringify({
          name: param.option?.label ?? "",
        }),
      });

      if (response.ok) {
        const category = (await response.json()) as Category;
        categoryOption = {
          id: category.id,
          label: category.name,
          category,
        };
        categories.push(categoryOption);
      }
    }

    if (categoryOption) {
      setCategory([categoryOption]);
    }
  };

  const onSubmit = async (): Promise<void> => {
    const response = await fetch("/api/v1/items", {
      method: "POST",
      body: JSON.stringify({
        name,
        quantity: quantity[0],
        categoryId: category[0]?.category.id,
        category: category[0]?.category,
      }),
    });

    if (response.ok) {
      const { id, category } = (await response.json()) as Item;
      items.push([id, name, quantity[0], category.name]);
      setName("");
      setQuantity([0]);
      setCategory([]);
    }
  };
  return (
    <div>
      <FormControl label={() => "Item name"}>
        <Input
          value={name}
          placeholder="Item name"
          onChange={(e) => setName(e.currentTarget.value)}
        />
      </FormControl>
      <FormControl label={() => "Quantity"}>
        <Slider
          value={quantity}
          onChange={({ value }) => value && setQuantity(value)}
        />
      </FormControl>
      <FormControl label={() => "Category"}>
        <Select
          size={SelectSize.compact}
          options={categories}
          value={category}
          type={SelectType.search}
          placeholder="Select category"
          onChange={(params) => onSelect(params)}
          labelKey="label"
          valueKey="id"
          creatable
        />
      </FormControl>

      <Button
        onClick={onSubmit}
        size={ButtonSize.compact}
        shape={ButtonType.pill}
      >
        Add item
      </Button>
    </div>
  );
};
