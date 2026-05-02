"use client";

import { Search } from "lucide-react";
import { SelectControl } from "@/components/forms/select-control";
import { useScrollCollapse } from "@/hooks/use-scroll-collapse";

const communityFilters = [
  {
    label: "Community category",
    options: ["All Categories", "Academic", "Social", "Residence", "Market"],
    className: "lg:flex-1",
  },
  {
    label: "Sort communities",
    options: ["Most Active", "Newest", "Largest"],
    className: "lg:flex-1",
  },
];

export function CommunityFilterPanel() {
  const isExpanded = useScrollCollapse();

  return (
    <>
      <div
        className={`transition-[height] duration-300 ease-out ${
          isExpanded ? "h-[150px]" : "h-0"
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
                  placeholder="Find communities..."
                  className="w-full rounded-[1.35rem] border border-gray-100 bg-brand-cream py-3.5 pl-14 pr-5 text-base font-semibold text-gray-900 shadow-sm outline-none transition-all placeholder:text-gray-400 focus:bg-white focus:ring-2 focus:ring-brand-orange/20"
                />
              </div>
            </div>

            <div className="grid gap-3 border-t border-gray-100 pt-3 sm:grid-cols-2">
              {communityFilters.map((filter) => (
                <FilterField key={filter.label} label={filter.label}>
                  <SelectControl
                    label={filter.label}
                    options={filter.options}
                    className={filter.className}
                    variant="panel"
                  />
                </FilterField>
              ))}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

interface FilterFieldProps {
  label: string;
  children: React.ReactNode;
}

function FilterField({ label, children }: FilterFieldProps) {
  return (
    <div>
      <div className="mb-1.5 min-h-5 px-1">
        <span className="text-[11px] font-bold uppercase tracking-[0.22em] text-gray-400">
          {label}
        </span>
      </div>
      {children}
    </div>
  );
}
