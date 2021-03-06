import { SexType, University, RoleType } from "infra/type";

export type PaymentStatus = {
  userId: number;
  authorizer: number;
  period: number;
  createdAt: string;
  updatedAt: string;
};

export type User = {
  id: number;
  email: string;
  generation: number;
  name: string;
  kana: string;
  handle: string;
  sex: SexType;
  university: University;
  studentId: string;
  emergencyPhoneNumber: string;
  otherCircles: string;
  workshops: string[];
  squads: string[];
  role: RoleType;

  slackId: string;
  discordId: string;
  slackInvitationStatus: string;

  createdAt: string;
  updatedAt: string;
  emailVerified: boolean;

  paymentStatus?: PaymentStatus;
};

export type UpdateUserProfileInput = Omit<
  User,
  | "slackId"
  | "role"
  | "slackInvitationStatus"
  | "createdAt"
  | "updatedAt"
  | "emailVerified"
  | "id"
>;
