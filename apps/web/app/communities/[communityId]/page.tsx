import { AvailableSpaces } from "@/components/community/available-spaces";
import { CommunityCTA } from "@/components/community/community-cta";
import { CommunityFooter } from "@/components/community/community-footer";
import { CommunityHeader } from "@/components/community/community-header";
import { CommunityHero } from "@/components/community/community-hero";

export default function CommunityPage() {
  return (
    <div className="min-h-screen bg-background">
      <CommunityHeader />
      <main>
        <CommunityHero />
        <AvailableSpaces />
        <CommunityCTA />
      </main>
      <CommunityFooter />
    </div>
  );
}
