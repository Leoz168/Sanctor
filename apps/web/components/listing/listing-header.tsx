"use client";

import Link from "next/link";
import { Bell, HelpCircle, Search } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export function ListingHeader() {
  return (
    <header className="sticky top-0 z-50 w-full border-b bg-card/95 backdrop-blur supports-[backdrop-filter]:bg-card/80">
      <div className="container mx-auto flex h-16 items-center justify-between px-4">
        <div className="flex items-center gap-8">
          <Link href="/" className="text-xl font-bold text-primary">
            Sanctor
          </Link>
          <nav className="hidden items-center gap-6 md:flex">
            <Link href="/" className="text-sm font-medium text-primary hover:text-primary/80">
              Discover
            </Link>
            <Link href="#" className="text-sm font-medium text-muted-foreground hover:text-foreground">
              Communities
            </Link>
            <Link href="#" className="text-sm font-medium text-muted-foreground hover:text-foreground">
              Listings
            </Link>
            <Link href="#" className="text-sm font-medium text-muted-foreground hover:text-foreground">
              Saved
            </Link>
          </nav>
        </div>

        <div className="flex items-center gap-4">
          <div className="relative hidden md:block">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              type="search"
              placeholder="Search communities or listings..."
              className="w-64 bg-secondary/50 pl-10"
            />
          </div>
          <Button variant="ghost" size="icon" className="text-muted-foreground">
            <Bell className="h-5 w-5" />
            <span className="sr-only">Notifications</span>
          </Button>
          <Button variant="ghost" size="icon" className="text-muted-foreground">
            <HelpCircle className="h-5 w-5" />
            <span className="sr-only">Help</span>
          </Button>
          <Button className="bg-primary text-primary-foreground hover:bg-primary/90">
            Message Poster
          </Button>
        </div>
      </div>
    </header>
  );
}
