"use client";

import { useState } from "react";
import { Search } from "lucide-react";
import { SelectControl } from "@/components/forms/select-control";
import { useScrollCollapse } from "@/hooks/use-scroll-collapse";

const housingFilters = [
  {
    label: "Start Term",
    options: ["Any Term", "Fall", "Spring", "Winter"],
    className: "lg:flex-1",
  },
  {
    label: "Rooms",
    options: ["Any Rooms", "1", "2", "3", "4", "5"],
    className: "lg:flex-1",
  },
  {
    label: "Gender Preference",
    options: ["Any Gender", "Coed", "Female Only", "Male Only"],
    className: "lg:flex-1",
  },
  {
    label: "Sort By",
    options: ["Most Recent", "Price: Low to High", "Price: High to Low"],
    className: "lg:flex-1",
  },
];

export function HousingFilterPanel() {
  const [maxPrice, setMaxPrice] = useState(3000);
  const isExpanded = useScrollCollapse();

  return (
    <>
      <div
        className={`transition-[height] duration-300 ease-out ${
          isExpanded ? "h-[178px]" : "h-0"
        }`}
      />

      <div
        className={`fixed left-0 right-0 top-20 z-30 border-t border-gray-100 bg-white/95 shadow-xl shadow-gray-900/5 backdrop-blur-md transition-all duration-300 ease-out ${
          isExpanded
            ? "translate-y-0 opacity-100"
            : "pointer-events-none -translate-y-full opacity-0"
        }`}
      >
        <div className="mx-auto max-w-7xl px-4 py-3 sm:px-6 lg:px-8">
          <div className="space-y-3 rounded-b-[1.5rem] border border-t-0 border-gray-100 bg-white/95 p-4 shadow-sm">
          <div className="flex flex-col gap-3 lg:flex-row lg:items-center">
            <div className="relative flex-1">
              <Search className="absolute left-5 top-1/2 h-5 w-5 -translate-y-1/2 text-gray-400" />
              <input
                type="text"
                placeholder="Search by neighborhood or building..."
                className="w-full rounded-[1.35rem] border border-gray-100 bg-brand-cream py-3.5 pl-14 pr-5 text-base font-semibold text-gray-900 shadow-sm outline-none transition-all placeholder:text-gray-400 focus:bg-white focus:ring-2 focus:ring-brand-orange/20"
              />
            </div>
          </div>

            <div className="grid gap-3 border-t border-gray-100 pt-3 lg:grid-cols-[1fr_1fr_1fr_1.05fr_1fr] lg:items-end">
              {housingFilters.slice(0, 3).map((filter) => (
                <FilterField key={filter.label} label={filter.label}>
                  <SelectControl
                    label={filter.label}
                    options={filter.options}
                    className={filter.className}
                    variant="panel"
                  />
                </FilterField>
              ))}

              <FilterField label="Max Price" value={`$${maxPrice}`}>
                <input
                  type="range"
                  min={500}
                  max={3000}
                  step={50}
                  value={maxPrice}
                  onChange={(event) => setMaxPrice(Number(event.target.value))}
                  className="h-9 w-full accent-brand-orange"
                />
              </FilterField>

              <FilterField label="Sort By">
                <SelectControl
                  label="Sort By"
                  options={housingFilters[3].options}
                  className={housingFilters[3].className}
                  variant="panel"
                />
              </FilterField>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

interface FilterFieldProps {
  label: string;
  value?: string;
  children: React.ReactNode;
}

function FilterField({ label, value, children }: FilterFieldProps) {
  return (
    <div>
      <div className="mb-1.5 flex min-h-5 items-center justify-between gap-3 px-1">
        <span className="text-[11px] font-bold uppercase tracking-[0.22em] text-gray-400">
          {label}
        </span>
        {value && (
          <span className="text-sm font-bold text-brand-orange">
            {value}
          </span>
        )}
      </div>
      {children}
    </div>
  );
}
