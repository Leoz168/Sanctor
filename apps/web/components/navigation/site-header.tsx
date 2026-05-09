"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { LogIn, Menu, User, X } from "lucide-react";
import { useEffect, useState } from "react";
import { BrandLogo } from "@/components/navigation/brand-logo";
import { getCurrentUser, type CurrentUser } from "@/lib/auth-client";

const navItems = [
  { href: "/post-listings", label: "Listings" },
  { href: "/communities", label: "Communities" },
];

export function SiteHeader() {
  const pathname = usePathname();
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [currentUser, setCurrentUser] = useState<CurrentUser | null>(null);

  useEffect(() => {
    let isMounted = true;

    getCurrentUser()
      .then((user) => {
        if (isMounted) {
          setCurrentUser(user);
        }
      })
      .catch(() => {
        if (isMounted) {
          setCurrentUser(null);
        }
      });

    return () => {
      isMounted = false;
    };
  }, [pathname]);

  const linkClassName = (href: string) =>
    `text-sm font-medium transition-colors ${
      pathname === href ? "text-brand-orange" : "text-gray-600 hover:text-brand-orange"
    }`;

  return (
    <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-gray-200">
      <div className="px-4 sm:px-6 lg:px-10">
        <div className="flex h-20 items-center justify-between">
          <BrandLogo />

          <div className="hidden md:flex items-center gap-8">
            {navItems.map((item) => (
              <Link key={item.href} href={item.href} className={linkClassName(item.href)}>
                {item.label}
              </Link>
            ))}
            <div className="h-4 w-px bg-gray-200" />
            <AuthNavLink currentUser={currentUser} />
          </div>

          <button
            onClick={() => setIsMenuOpen((open) => !open)}
            className="md:hidden p-2 text-gray-600 hover:text-brand-orange transition-colors"
            id="mobile-menu-toggle"
            aria-label="Toggle menu"
          >
            {isMenuOpen ? <X size={24} /> : <Menu size={24} />}
          </button>
        </div>
      </div>

      {isMenuOpen && (
        <div className="md:hidden bg-white border-b border-gray-200 p-4 space-y-4">
          {navItems.map((item) => (
            <Link
              key={item.href}
              href={item.href}
              onClick={() => setIsMenuOpen(false)}
              className="block px-3 py-2 text-base font-medium text-gray-700 hover:text-brand-orange"
            >
              {item.label}
            </Link>
          ))}
          <AuthNavLink currentUser={currentUser} onClick={() => setIsMenuOpen(false)} isMobile />
        </div>
      )}
    </nav>
  );
}

interface AuthNavLinkProps {
  currentUser: CurrentUser | null;
  onClick?: () => void;
  isMobile?: boolean;
}

function AuthNavLink({ currentUser, onClick, isMobile = false }: AuthNavLinkProps) {
  const href = currentUser ? "/profile" : "/login";
  const label = currentUser ? currentUser.username : "Login";
  const Icon = currentUser ? User : LogIn;
  const className = isMobile
    ? "w-full flex items-center justify-center gap-2 px-5 py-3 bg-brand-orange text-white rounded-full font-medium"
    : "flex items-center gap-2 px-5 py-2.5 bg-brand-orange text-white rounded-full font-medium shadow-md hover:shadow-lg hover:bg-orange-600 transition-all active:scale-95";

  return (
    <Link href={href} onClick={onClick} className={className} id={currentUser ? "nav-profile" : "nav-login"}>
      <Icon size={18} />
      <span>{label}</span>
    </Link>
  );
}
