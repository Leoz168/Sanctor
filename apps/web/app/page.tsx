import { AppShell } from "@/components/layout/app-shell";
import { FeatureSection } from "@/components/home/feature-section";
import { HomeHero } from "@/components/home/home-hero";
import { PostListingCta } from "@/components/home/post-listing-cta";

export default function HomePage() {
  return (
    <AppShell surface="cream">
      <HomeHero />
      <FeatureSection />
      <PostListingCta />
    </AppShell>
  );
}
