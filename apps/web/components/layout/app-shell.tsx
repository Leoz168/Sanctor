import { ReactNode } from "react";
import { SiteFooter } from "@/components/navigation/site-footer";
import { SiteHeader } from "@/components/navigation/site-header";

interface AppShellProps {
  children: ReactNode;
  floatingAction?: ReactNode;
  surface?: "white" | "cream";
}

export function AppShell({ children, floatingAction, surface = "white" }: AppShellProps) {
  const surfaceClassName = surface === "cream" ? "bg-brand-cream" : "bg-white";

  return (
    <div className={`min-h-screen flex flex-col ${surfaceClassName} font-sans`}>
      <SiteHeader />
      <main className="flex-1">{children}</main>
      <SiteFooter />
      {floatingAction}
    </div>
  );
}
