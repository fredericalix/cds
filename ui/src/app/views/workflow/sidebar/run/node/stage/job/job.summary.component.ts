import {Component, Input, OnInit} from '@angular/core';
import {Router, ActivatedRoute} from '@angular/router';
import {PipelineStatus} from '../../../../../../../model/pipeline.model';
import {WorkflowRun, WorkflowNodeJobRun} from '../../../../../../../model/workflow.run.model';
import {Stage} from '../../../../../../../model/stage.model';

@Component({
    selector: 'app-job-step-summary',
    templateUrl: './job.summary.component.html',
    styleUrls: ['./job.summary.component.scss']
})
export class JobStepSummaryComponent implements OnInit {

    @Input() job: WorkflowNodeJobRun;
    @Input() workflowRun: WorkflowRun;
    @Input() stage: Stage;

    open = false;
    constructor(private _router: Router, private _route: ActivatedRoute) {

    }

    ngOnInit() {
      this.open = this.job.status === PipelineStatus.FAIL || PipelineStatus.isActive(this.job.status);
    }

    goToJobLogs() {
      this._router.navigate([
          'project',
          this._route.snapshot.params['key'],
          'workflow',
          this._route.snapshot.params['workflowName'],
          'run',
          this.workflowRun.num,
          'node',
          this.job.workflow_node_run_id
      ], {queryParams: {stageId: this.stage.id, actionId: this.job.job.pipeline_action_id}});
    }
}