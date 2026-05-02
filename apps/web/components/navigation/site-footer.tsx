import Link from "next/link";
import { BrandLogo } from "@/components/navigation/brand-logo";

const footerLinks = ["Terms", "Privacy", "Contact", "Help"];

export function SiteFooter() {
  return (
    <footer className="bg-gray-50 border-t border-gray-200 pt-16 pb-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col md:flex-row justify-between items-center gap-8 mb-12">
          <BrandLogo size="sm" />
          <div className="flex gap-8 text-sm font-medium text-gray-500">
            {footerLinks.map((label) => (
              <Link key={label} href="#" className="hover:text-brand-orange">
                {label}
              </Link>
            ))}
          </div>
        </div>
        <div className="text-center text-gray-400 text-xs font-medium">
          &copy; {new Date().getFullYear()} Renting. Dedicated to student housing solutions.
        </div>
      </div>
    </footer>
  );
}
