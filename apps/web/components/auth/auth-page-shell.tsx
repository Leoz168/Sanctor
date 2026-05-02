import { ReactNode } from "react";
import { BrandLogo } from "@/components/navigation/brand-logo";

interface AuthPageShellProps {
  children: ReactNode;
}

export function AuthPageShell({ children }: AuthPageShellProps) {
  return (
    <main className="min-h-screen bg-brand-cream px-4 py-10 font-sans text-[#1A1A1A]">
      <div className="mx-auto flex min-h-[calc(100vh-5rem)] w-full max-w-6xl flex-col">
        <nav className="flex items-center justify-between">
          <BrandLogo />
        </nav>

        <section className="flex flex-1 items-center justify-center py-12">
          {children}
        </section>
      </div>
    </main>
  );
}
