//@ts-nocheck
import { lockQueryTypeFromJSON } from "../lockup/lock";
import { Long } from "../../helpers";
export const AminoConverter = {
  "/osmosis.incentives.MsgCreateGauge": {
    aminoType: "osmosis/incentives/create-gauge",
    toAmino: ({
      isPerpetual,
      owner,
      distributeTo,
      coins,
      startTime,
      numEpochsPaidOver
    }) => {
      return {
        is_perpetual: isPerpetual,
        owner,
        distribute_to: {
          lock_query_type: distributeTo.lockQueryType,
          denom: distributeTo.denom,
          duration: (distributeTo.duration * 1000000000).toString(),
          timestamp: distributeTo.timestamp
        },
        coins: coins.map(el0 => ({
          denom: el0.denom,
          amount: el0.amount
        })),
        start_time: startTime,
        num_epochs_paid_over: numEpochsPaidOver.toString()
      };
    },
    fromAmino: ({
      is_perpetual,
      owner,
      distribute_to,
      coins,
      start_time,
      num_epochs_paid_over
    }) => {
      return {
        isPerpetual: is_perpetual,
        owner,
        distributeTo: {
          lockQueryType: lockQueryTypeFromJSON(distribute_to.lock_query_type),
          denom: distribute_to.denom,
          duration: {
            seconds: Long.fromNumber(Math.floor(parseInt(distribute_to.duration) / 1000000000)),
            nanos: parseInt(distribute_to.duration) % 1000000000
          },
          timestamp: distribute_to.timestamp
        },
        coins: coins.map(el0 => ({
          denom: el0.denom,
          amount: el0.amount
        })),
        startTime: start_time,
        numEpochsPaidOver: Long.fromString(num_epochs_paid_over)
      };
    }
  },
  "/osmosis.incentives.MsgAddToGauge": {
    aminoType: "osmosis/incentives/add-to-gauge",
    toAmino: ({
      owner,
      gaugeId,
      rewards
    }) => {
      return {
        owner,
        gauge_id: gaugeId.toString(),
        rewards: rewards.map(el0 => ({
          denom: el0.denom,
          amount: el0.amount
        }))
      };
    },
    fromAmino: ({
      owner,
      gauge_id,
      rewards
    }) => {
      return {
        owner,
        gaugeId: Long.fromString(gauge_id),
        rewards: rewards.map(el0 => ({
          denom: el0.denom,
          amount: el0.amount
        }))
      };
    }
  }
};