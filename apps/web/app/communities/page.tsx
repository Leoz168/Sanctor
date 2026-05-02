import { MessageSquare } from "lucide-react";
import { CommunityFilterPanel } from "@/components/catalog/community-filter-panel";
import { FloatingActionButton } from "@/components/catalog/floating-action-button";
import { CommunityCard } from "@/components/community-card";
import { AppShell } from "@/components/layout/app-shell";

const communities = [
  {
    id: 1,
    href: "/communities/computer-science-collective",
    name: "Computer Science Collective",
    description: "The primary hub for CSS/CS students to share resources, study groups, and tech news.",
    category: "ACADEMIC",
    members: 1248,
    postsPerWeek: 382,
    image: "/images/community-1.jpg",
    isJoined: true,
  },
  {
    id: 2,
    href: "/communities/residence-life-hub",
    name: "Residence Life Hub",
    description: "Stay updated on campus events, cafeteria menus, and dorm-specific announcements.",
    category: "RESIDENCE",
    members: 1892,
    postsPerWeek: 524,
    isJoined: false,
  },
  {
    id: 3,
    href: "/communities/student-market-swap",
    name: "Student Market & Swap",
    description: "Buy and sell textbooks, electronics, and furniture within the university community.",
    category: "MARKET",
    members: 3450,
    postsPerWeek: 890,
    isJoined: false,
  },
  {
    id: 4,
    href: "/communities/campus-social-circle",
    name: "Campus Social Circle",
    description: "Casual discussion for weekend plans, club events, and general campus life.",
    category: "SOCIAL",
    members: 2100,
    postsPerWeek: 412,
    image: "/images/community-4.jpg",
    isJoined: false,
  },
  {
    id: 5,
    href: "/communities/art-studio-zone",
    name: "The Art Studio Zone",
    description: "A creative space for visual artists, photographers, and designers to collaborate.",
    category: "SOCIAL",
    members: 850,
    postsPerWeek: 125,
    image: "/images/community-5.jpg",
    isJoined: false,
  },
  {
    id: 6,
    href: "/communities/pre-med-study-group",
    name: "Pre-Med Study Group",
    description: "Advanced study materials, MCAT prep, and lab report discussions.",
    category: "ACADEMIC",
    members: 1100,
    postsPerWeek: 245,
    isJoined: false,
  },
];

export default function CommunitiesPage() {
  return (
    <AppShell floatingAction={<FloatingActionButton icon={MessageSquare}>Create a community</FloatingActionButton>}>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <CommunityFilterPanel />

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {communities.map((community) => (
            <CommunityCard key={community.id} {...community} />
          ))}
        </div>
      </div>
    </AppShell>
  );
}
