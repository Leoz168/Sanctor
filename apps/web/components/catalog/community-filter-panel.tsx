"use client";

import { MessageSquare, Search } from "lucide-react";
import { AutoCollapsingFilterShell } from "@/components/catalog/auto-collapsing-filter-shell";
import { SelectControl } from "@/components/forms/select-control";

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

interface CommunityFilterPanelProps {
  onCreateCommunity: () => void;
}

export function CommunityFilterPanel({
  onCreateCommunity,
}: CommunityFilterPanelProps) {
  return (
    <AutoCollapsingFilterShell
      expandedHeightClassName="h-[318px] sm:h-[220px] lg:h-[150px]"
      action={
        <button
          type="button"
          onClick={onCreateCommunity}
          className="inline-flex items-center gap-2 rounded-full bg-brand-orange px-5 py-3 text-sm font-bold text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-95"
        >
          <MessageSquare className="h-5 w-5" />
          <span>Create community</span>
        </button>
      }
    >
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
    </AutoCollapsingFilterShell>
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
