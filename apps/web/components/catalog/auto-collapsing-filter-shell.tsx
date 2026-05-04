"use client";

import { ReactNode } from "react";
import { useScrollCollapse } from "@/hooks/use-scroll-collapse";

interface AutoCollapsingFilterShellProps {
  children: ReactNode;
  expandedHeightClassName: string;
  action?: ReactNode;
}

export function AutoCollapsingFilterShell({
  children,
  expandedHeightClassName,
  action,
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
        className={`fixed left-0 right-0 top-20 z-30 border-b border-gray-100 bg-white/95 shadow-md shadow-gray-900/5 backdrop-blur-md transition-all duration-300 ease-out ${
          isExpanded
            ? "translate-y-0 opacity-100"
            : "pointer-events-none -translate-y-full opacity-0"
        }`}
      >
        <div className="mx-auto flex max-w-[92rem] items-start gap-4 px-4 py-4 sm:px-6 lg:px-8">
          <div className="min-w-0 flex-1 space-y-3">
            {children}
          </div>
          {action && <div className="shrink-0 pt-0.5">{action}</div>}
        </div>
      </div>
    </>
  );
}
