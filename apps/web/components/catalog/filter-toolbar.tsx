import { ReactNode } from "react";
import { SlidersHorizontal } from "lucide-react";
import { SearchInput } from "@/components/forms/search-input";
import { SelectControl } from "@/components/forms/select-control";

interface FilterToolbarProps {
  searchPlaceholder: string;
  filters: Array<{
    label: string;
    options: string[];
    className?: string;
  }>;
  actionIcon?: ReactNode;
}

export function FilterToolbar({ searchPlaceholder, filters, actionIcon }: FilterToolbarProps) {
  return (
    <div className="flex flex-col lg:flex-row lg:items-center gap-4 mb-10">
      <SearchInput placeholder={searchPlaceholder} />

      <div className="flex flex-col sm:flex-row gap-3">
        {filters.map((filter) => (
          <SelectControl
            key={filter.label}
            label={filter.label}
            options={filter.options}
            className={filter.className}
          />
        ))}

        <button className="h-[54px] w-full sm:w-[54px] rounded-2xl bg-brand-cream border border-gray-100 text-gray-700 hover:bg-white transition-colors flex items-center justify-center">
          {actionIcon ?? <SlidersHorizontal size={20} />}
          <span className="sr-only">More filters</span>
        </button>
      </div>
    </div>
  );
}
