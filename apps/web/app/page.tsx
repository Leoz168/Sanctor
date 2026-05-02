"use client";

import Link from "next/link";
import { useState } from "react";
import {
  ChevronDown,
  Home,
  Mail,
  Menu,
  MessageSquare,
  Search,
  ShoppingBag,
  User,
  X,
} from "lucide-react";
import { AppShell } from "@/components/layout/app-shell";
import { FeatureSection } from "@/components/home/feature-section";
import { HomeHero } from "@/components/home/home-hero";

export default function HomePage() {
  return (
    <AppShell surface="cream">
      <HomeHero />
      <FeatureSection />
    </AppShell>
  );
}
