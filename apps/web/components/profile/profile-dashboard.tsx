"use client";

import { useState } from "react";
import { Star } from "lucide-react";
import { ProfileBookmarksPanel } from "@/components/profile/profile-bookmarks-panel";
import { ProfileCommunitiesPanel } from "@/components/profile/profile-communities-panel";
import { ProfileFormCard } from "@/components/profile/profile-form-card";
import { ProfileListingsPanel } from "@/components/profile/profile-listings-panel";
import { ProfilePlaceholderPanel } from "@/components/profile/profile-placeholder-panel";
import {
  ProfileSidebar,
  type ProfileTab,
} from "@/components/profile/profile-sidebar";

export function ProfileDashboard() {
  const [activeTab, setActiveTab] = useState<ProfileTab>("profile");

  return (
    <>
      <ProfileSidebar activeTab={activeTab} onTabChange={setActiveTab} />
      {activeTab === "profile" ? <ProfileFormCard /> : null}
      {activeTab === "listings" ? <ProfileListingsPanel /> : null}
      {activeTab === "communities" ? <ProfileCommunitiesPanel /> : null}
      {activeTab === "bookmarks" ? <ProfileBookmarksPanel /> : null}
      {activeTab === "account" ? (
        <ProfilePlaceholderPanel
          title="My Account"
          description="Account, notification, and privacy settings will be available here."
          icon={<Star className="h-6 w-6" />}
        />
      ) : null}
    </>
  );
}
