import Link from "next/link";
import { LucideIcon } from "lucide-react";

interface FloatingActionButtonProps {
  icon: LucideIcon;
  children: string;
  href?: string;
  placement?: "fixed" | "inline" | "top-right";
}

export function FloatingActionButton({
  icon: Icon,
  children,
  href,
  placement = "fixed",
}: FloatingActionButtonProps) {
  const content = (
    <>
      <Icon className="w-5 h-5" />
      <span>{children}</span>
    </>
  );
  const className =
    placement === "fixed"
      ? "fixed bottom-8 right-6 sm:right-8 flex items-center gap-2 px-5 py-4 bg-brand-orange text-white rounded-full font-bold shadow-xl shadow-brand-orange/30 hover:bg-orange-600 transition-all active:scale-95"
      : placement === "top-right"
        ? "fixed right-4 top-24 z-40 inline-flex items-center gap-2 rounded-full bg-brand-orange px-5 py-3 text-sm font-bold text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-95 sm:right-8 lg:right-12"
      : "inline-flex items-center gap-2 rounded-full bg-brand-orange px-5 py-3 text-sm font-bold text-white shadow-lg shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-95";

  return href ? (
    <Link href={href} className={className}>
      {content}
    </Link>
  ) : (
    <button className={className}>
      {content}
    </button>
  );
}
