import Link from "next/link";
import { Home } from "lucide-react";

interface BrandLogoProps {
  size?: "sm" | "md";
}

export function BrandLogo({ size = "md" }: BrandLogoProps) {
  const iconSize = size === "sm" ? 18 : 24;
  const iconPadding = size === "sm" ? "p-1.5 rounded-lg" : "p-2 rounded-xl";
  const textSize = size === "sm" ? "text-xl" : "text-2xl";

  return (
    <Link href="/" className="flex items-center gap-2 group" id="nav-logo">
      <div className={`${iconPadding} bg-brand-orange text-white shadow-lg shadow-brand-orange/20`}>
        <Home size={iconSize} />
      </div>
      <span className={`${textSize} font-bold tracking-tight text-gray-900 group-hover:text-brand-orange transition-colors`}>
        Rentling
      </span>
    </Link>
  );
}
