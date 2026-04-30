"use client";

import { useState } from "react";
import Link from "next/link";
import { Home, Menu, X, ChevronDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

export function Header() {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  return (
    <header className="sticky top-0 z-50 w-full border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
        <Link href="/" className="flex items-center gap-2">
          <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-accent">
            <Home className="h-5 w-5 text-accent-foreground" />
          </div>
          <span className="text-xl font-semibold text-foreground">Sanctor</span>
        </Link>

        <nav className="hidden items-center gap-8 md:flex">
          <DropdownMenu>
            <DropdownMenuTrigger className="flex items-center gap-1 text-sm text-muted-foreground transition-colors hover:text-foreground">
              Discover <ChevronDown className="h-4 w-4" />
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>Academic communities</DropdownMenuItem>
              <DropdownMenuItem>Residence groups</DropdownMenuItem>
              <DropdownMenuItem>Creator circles</DropdownMenuItem>
              <DropdownMenuItem>Support networks</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>

          <DropdownMenu>
            <DropdownMenuTrigger className="flex items-center gap-1 text-sm text-muted-foreground transition-colors hover:text-foreground">
              Tools <ChevronDown className="h-4 w-4" />
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>Direct messaging</DropdownMenuItem>
              <DropdownMenuItem>Role management</DropdownMenuItem>
              <DropdownMenuItem>Safety controls</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>

          <Link href="#" className="text-sm text-muted-foreground transition-colors hover:text-foreground">
            Campuses
          </Link>
          <Link href="#" className="text-sm text-muted-foreground transition-colors hover:text-foreground">
            Moderation
          </Link>
        </nav>

        <div className="hidden items-center gap-4 md:flex">
          <Button variant="ghost" size="sm" className="text-muted-foreground">
            Log in
          </Button>
          <Button size="sm" className="bg-primary text-primary-foreground hover:bg-primary/90">
            Launch Sanctor
          </Button>
        </div>

        <button
          className="md:hidden"
          onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          aria-label="Toggle menu"
        >
          {mobileMenuOpen ? (
            <X className="h-6 w-6 text-foreground" />
          ) : (
            <Menu className="h-6 w-6 text-foreground" />
          )}
        </button>
      </div>

      {mobileMenuOpen && (
        <div className="border-t border-border bg-background md:hidden">
          <div className="space-y-1 px-4 py-4">
            <Link href="#" className="block py-2 text-foreground">Discover</Link>
            <Link href="#" className="block py-2 text-foreground">Tools</Link>
            <Link href="#" className="block py-2 text-foreground">Campuses</Link>
            <Link href="#" className="block py-2 text-foreground">Moderation</Link>
            <div className="flex flex-col gap-2 pt-4">
              <Button variant="ghost" className="w-full justify-start text-muted-foreground">
                Log in
              </Button>
              <Button className="w-full bg-primary text-primary-foreground">
                Launch Sanctor
              </Button>
            </div>
          </div>
        </div>
      )}
    </header>
  );
}
