"use client";

import { ReactNode } from "react";
import { useScrollCollapse } from "@/hooks/use-scroll-collapse";

interface AutoCollapsingFilterShellProps {
  children: ReactNode;
  expandedHeightClassName: string;
}

export function AutoCollapsingFilterShell({
  children,
  expandedHeightClassName,
}: AutoCollapsingFilterShellProps) {
  const isExpanded = useScrollCollapse();

  return (
    <>
      <div
        className={`transition-[height] duration-300 ease-out ${
          isExpanded ? expandedHeightClassName : "h-0"
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
            {children}
          </div>
        </div>
      </div>
    </>
  );
}
